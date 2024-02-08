import binascii
from web3 import Web3
w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))

with open("./UTC--2022-08-19T17-38-31.257380510Z--123463a4b065722e99115d6c222f267d9cabb524") as keyfile:
    encrypted_key = keyfile.read()
    private_key = w3.eth.account.decrypt(encrypted_key, '')
# print(private_key)
# print(binascii.b2a_hex(private_key))
print("Private key:",binascii.b2a_hex(private_key).decode())
# print("0x"+binascii.b2a_hex(private_key).decode())