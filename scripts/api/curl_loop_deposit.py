import requests
import threading

count_iters_for_thread = 10
count_threads = 10
value_amount_deposit = 1


headers = {
	'Content-Type': 'application/json',
}

data_deposit = '{"walletId":"acd375bf-0f3e-4877-a993-d1c0733096af", "amount":'+str(value_amount_deposit)+', "operationType":"DEPOSIT"}'


def request_deposit():
	response = requests.post('http://127.0.0.1:8080/api/v1/wallet', headers=headers, data=data_deposit)

def request_balance():
	response = requests.get('http://127.0.0.1:8080/api/v1/wallets/acd375bf-0f3e-4877-a993-d1c0733096af')
	return response.json()['Data']['balance']

def request_test():
	for num_thread in range(count_iters_for_thread):
		# print(f"{num_thread} -- Thread")
		request_balance()
		request_deposit()

value_start_balance = float(request_balance())
value_wait_end_balance = count_iters_for_thread * count_threads * value_amount_deposit + value_start_balance

print("starting test api")

for i in range(count_threads):
	th = threading.Thread(target=request_test)
	th.start()
	th.join()


value_end_balance = float(request_balance())

print(f"balance start: {value_start_balance}")
print(f"balance end: {value_end_balance}")
print(f"balance waiting: {value_wait_end_balance}")