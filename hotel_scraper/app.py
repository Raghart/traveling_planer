#FastAPI App
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
#import uvicorn
import requests
import json
from playwright.sync_api import sync_playwright
import re
import sqlite3
from datetime import datetime, timedelta, timezone
from dotenv import load_dotenv
import os

load_dotenv()
GEOAPIFY_KEY = os.getenv("GEOAPIFY_KEY")

app = FastAPI()

DB_PATH = "scraper_cache.db"
CACHE_EXPIRATION_DAYS = 30  # 1 month

def get_cached_data(country_code: str):
    """Retrieves data from SQLite."""
    with sqlite3.connect(DB_PATH) as conn:
        cursor = conn.cursor()
        try:
            cursor.execute(
                "SELECT scraped_data, last_updated FROM country_cache WHERE country_code = ?", 
                (country_code,)
            )
        except sqlite3.OperationalError as e:
            if "no such table" in str(e):
                # Table doesn't exist, create it and return None
                cursor.execute(
                    """
                    CREATE TABLE IF NOT EXISTS country_cache (
                        country_code TEXT PRIMARY KEY,
                        scraped_data TEXT,
                        last_updated TEXT
                    )
                    """
                )
                conn.commit()
                return None
            else:
                raise e

        row = cursor.fetchone()
        if row:
            return row

def update_cache(country_code: str, data: dict):
    """Saves or updates the scraped data in SQLite."""
    with sqlite3.connect(DB_PATH) as conn:
        cursor = conn.cursor()
        cursor.execute(
            """
            INSERT OR REPLACE INTO country_cache (country_code, scraped_data, last_updated)
            VALUES (?, ?, ?)
            """,
            (country_code, json.dumps(data), datetime.now(timezone.utc).isoformat())
        )
        conn.commit()

def obtain_coordinates(address):
    # Endpoint for Nominatim API
    url = f"https://nominatim.openstreetmap.org/search"
    
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0'
    }
    
    params = {
        'q': address,
        'format': 'json',
        'limit': 1
    }
    try:
        response = requests.get(url, headers=headers, params=params)
    except requests.RequestException as e:
        print(f"Error fetching coordinates: {e}")
        return None

    if response.status_code == 200 and len(response.json()) > 0:
        data = response.json()[0]
        bbox = data['boundingbox']
        
        sw_corner = (float(bbox[0]), float(bbox[2])) # Lat, Lon
        ne_corner = (float(bbox[1]), float(bbox[3])) # Lat, Lon
        
        return {
            "southwest": sw_corner,
            "northeast": ne_corner
        }
    else:
        print(f"Failed to obtain coordinates for address: {address}. Status code: {response.status_code}")
        return None

def obtain_hotels(query: str = "Caracas, Venezuela"):
    categories = "accommodation.hotel"
    coordinates = obtain_coordinates(query)
    if not coordinates:
        print("Failed to obtain coordinates for the given query.")
        return None
    filter = f"rect:{coordinates['southwest'][1]},{coordinates['southwest'][0]},{coordinates['northeast'][1]},{coordinates['northeast'][0]}"
    limit = 10
    apiKey = GEOAPIFY_KEY
    url = f"https://api.geoapify.com/v2/places?categories={categories}&filter={filter}&limit={limit}&apiKey={apiKey}"

    try:
        response = requests.get(url)
    except requests.RequestException as e:
        print(f"Error fetching hotel data: {e}")
        return None

    data = response.json()
    print("Parsed JSON data:")
    print(json.dumps(data, indent=4))

    return data

def scrape_ratings(query:str):
    with sync_playwright() as p:
        response = obtain_hotels(query)
        if not response:
            print("Failed to obtain hotel data.")
            return None

        browser = p.firefox.launch(headless=True)
        context = browser.new_context(
            viewport={'width': 1280, 'height': 1280},
            locale='en-US',
            extra_http_headers={"Accept-Language": "en-US,en;q=0.9"}
        )

        page = context.new_page()
        hotel_data = []
            
        try:
            page.goto("https://www.google.com/maps?hl=en")

            for i in range(0, len(response['features'])):
                hotel_name = response['features'][i]['properties']['name']
                hotel_address = response['features'][i]['properties']['city']
                search_query = f"{hotel_name}, {hotel_address}"
                
                search_bar = page.locator('input.UGojuc') # input id= ucc-1
                search_bar.click(timeout=5000)
                #Delete content of search bar
                search_bar.fill("")
                search_bar.fill(search_query)
                search_bar.press("Enter")
                page.wait_for_timeout(5000)

                try:
                    rating =page.locator('div.F7nice').inner_text(timeout=5000)
                    rating_value = float(rating.split('\n')[0]) # Extract the numeric part
                    amount_of_reviews = int(re.findall(r'\d+', rating.split('\n')[1])[0]) # Extract the number of reviews
                    rating_score = rating_value * amount_of_reviews
                    
                    print(f"{hotel_name}: {rating_score}")
                    hotel_data.append({
                        "name": hotel_name,
                        "score": rating_score
                    })
                except:
                    print(f"{hotel_name}: No rating found")
        except:
            print(f"Failed to search for {hotel_name}")

        finally:
            browser.close()

    return json.dumps(hotel_data, indent=4)

class Item(BaseModel): #Received data model
    country: str
    city: str = None

@app.post("/hotels")
def get_hotels(item: Item):
    cache_record = get_cached_data(item.country)
    if cache_record:
        cached_data, last_updated = cache_record
        last_updated = datetime.fromisoformat(last_updated)
        data_age = datetime.now(timezone.utc) - last_updated
        if data_age < timedelta(days=CACHE_EXPIRATION_DAYS):
            print("Returning cached data")
            return json.loads(cached_data)

    try:
        query = f"{item.city}, {item.country}" if item.city else item.country
        scraped_data = scrape_ratings(query=query)
        if scraped_data: print("Scraped data successfully obtained")
        update_cache(item.country, json.loads(scraped_data))
        return json.loads(scraped_data)
    except Exception as e:
        raise HTTPException(status_code=500, detail="An error occurred while fetching hotel data: " + str(e))

    
    