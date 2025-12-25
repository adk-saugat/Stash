from db import get_users

users = get_users()

for user in users:
    print(user["name"])

