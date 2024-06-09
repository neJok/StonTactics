from pymongo import MongoClient

client = MongoClient("mongodb://Ewxq1EtAjP:kWLbt5lGpP@109.107.189.55:27017/")
db = client.mydbname.users

print(db.find_one({}))