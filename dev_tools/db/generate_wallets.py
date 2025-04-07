import psycopg2 as psql
from random import randint

DB_NAME = "wallets"
DB_USER = "postgres"
DB_PASS = "postgres9872"
DB_HOST = "127.0.0.1"
DB_PORT = "5432"

NUM_GENERATE_WALLETS = 10
MIN_VALUE_BALANCE = 0
MAX_VALUE_BALANCE = 9999

query_drop_all_wallets = f"DROP TABLE {DB_NAME};"
query_create_table_wallets = f"CREATE TABLE {DB_NAME} (id UUID NOT NULL, balance NUMERIC(10,2) NOT NULL);"
query_new_wallet = f"INSERT INTO {DB_NAME} (id, balance) VALUES (gen_random_uuid(), 0)"
query_new_wallet_from_balance = lambda balance : f"INSERT INTO {DB_NAME} (id, balance) VALUES (gen_random_uuid(), {balance})"
query_get_all_wallets = f"SELECT * FROM {DB_NAME};"

# !!! DROP ALL ROWS (ALL WALLETS) FROM TABLE !!!
def drop_all_wallets_from_database(cursor):
	cursor.execute(query_drop_all_wallets)	

def create_table_wallets(cursor):
	cursor.execute(query_create_table_wallets)	

def print_wallets(cursor):
	cursor.execute(query_get_all_wallets)
	i = 0
	for wallet in cursor.fetchall():
		print(f"Wallet:{i+1} {wallet}")
		i+=1

def new_wallet(cursor):
	cursor.execute(query_new_wallet)

def new_wallet_from_balance(cursor, balance):
	cursor.execute(query_new_wallet_from_balance(balance))

def generate_wallets_empty_balance():
	conn = psql.connect(dbname=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)
	conn.autocommit = True
	with conn:
		with conn.cursor() as cursor:
			drop_all_wallets_from_database(cursor)
			create_table_wallets(cursor)
			for _ in range(NUM_GENERATE_WALLETS):
				new_wallet(cursor)
			print_wallets(cursor)
	cursor.close()
	conn.close()

def generate_wallets_random_balance():
	conn = psql.connect(dbname=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)
	conn.autocommit = True
	with conn:
		with conn.cursor() as cursor:
			drop_all_wallets_from_database(cursor)
			create_table_wallets(cursor)
			for _ in range(NUM_GENERATE_WALLETS):
				new_wallet_from_balance(cursor, randint(MIN_VALUE_BALANCE, MAX_VALUE_BALANCE))
			print_wallets(cursor)
	cursor.close()
	conn.close()

def main():
	generate_wallets_random_balance()

main()