import asyncio
import websockets
import requests
import json
import base64
from dataclasses import dataclass

ClientDetails = {
    "Username": None,
    "Salt": None,
    "G": None,
    "K": None,
    "N": None,
    "V": None
}

ClientTempDetails = {
    "A": None,
    "B": None,
    "a": None,
    "u": None,
    "K_client": None
}

U_generation = {
    "A": None,
    "B": None
}

priv_vars = {
    "a":None,
    "u":None
}
def is_int(s):
    try:
        int(s)
        return True
    except ValueError:
        return False


def call_compute_a_api(base_url, data):
    try:
        response = requests.post(f"{base_url}", json=data)

        if response.status_code == 200:
            decoded_message = base64.b64decode(response.json()['message']).decode('utf-8')
            json_object = json.loads(decoded_message)
            return json_object

        elif response.status_code == 409:
            print('Signup Failure:', response.json())
            return {"error": "Signup Failure"}
        else:
            print('Error:', response.status_code)
            print(response.text)

    except requests.exceptions.RequestException as e:
        return {"error": f"Request error: {e}"}


async def login_101(uri, data):
    async with websockets.connect(uri) as ws:

        request_101 = json.dumps({"status": 101, "message": data})
        await ws.send(request_101)

        response_101 = await ws.recv()
        outer_dict = json.loads(response_101)
        inner_dict = json.loads(outer_dict['message'])
        json_data = json.loads(inner_dict['message'])

        return inner_dict, json_data

async def login_201(uri, data):
    async with websockets.connect(uri) as ws:

        request_201 = json.dumps({"status": 201, "message": data})
        await ws.send(request_201)
        response_201 = await ws.recv()
        response_201 = json.loads(response_201)
        
        return response_201['message']


async def login_301(uri):
    async with websockets.connect(uri) as ws:

        request_301 = json.dumps({"status": 301, "message": "Hello, server!"})
        await ws.send(request_301)

        response_301 = await ws.recv()
        response_301 = json.loads(response_301)
        return response_301['message']


async def login_401(uri, data):
    async with websockets.connect(uri) as ws:

        request_401 = json.dumps({"status": 401, "message": data})
        await ws.send(request_401)

        response_301 = await ws.recv()
        response_301 = json.loads(response_301)
        return response_301['message']


async def login_501(uri, data):
    async with websockets.connect(uri) as ws:

        request_501 = json.dumps({"status": 501, "message": data})
        await ws.send(request_501)

        response_501 = await ws.recv()
        response_501 = json.loads(response_501)
        return response_501['message']


async def main():
    uri = "ws://localhost:2002/ws" 
    Endpoint = "http://localhost:2004/"

    # Step 1: Handle 101
    name = input("Enter your name: ")
    password_input = input("Enter your password: ")
    if not is_int(password_input):
        password = password_input
    else:
        password = int(password_input)

    innerdict, jsondict = await login_101(uri, name)
    for key in jsondict:
        if key in ClientDetails:
            ClientDetails[key] = jsondict[key]
    ClientTempDetails['B'] = innerdict["metadata"]
    print(ClientTempDetails['B'])

    value_A_a = call_compute_a_api(f"{Endpoint}computeA", jsondict)

    ClientTempDetails['A'] = value_A_a["A"]
    ClientTempDetails['a'] = value_A_a["a"]

    print("A: ---> ",ClientTempDetails['A'], "B: ---> " ,ClientTempDetails['B'], "a: ---> ", ClientTempDetails['a'])
    print("Completion of 101")


    # Step 2: Handle 201
    computation_result = await login_201(uri, str(ClientTempDetails["A"]))
    if computation_result:
        U_generation["A"] = ClientTempDetails["A"]
        U_generation["B"] = ClientTempDetails["B"]
        value_U = call_compute_a_api(f"{Endpoint}computeU", U_generation)
        ClientTempDetails["u"] = value_U
        print("u: ---> ", ClientTempDetails["u"])
    print("Completion of 201")


    # Step 3: Handle 301
    computation_result_301 = await login_301(uri)
    data = {
        'user': ClientDetails,
        'user_tempdetails': ClientTempDetails,
        'priv_vars': {
            'a': ClientTempDetails['a'],
            'u': ClientTempDetails['u']
        },
        'U_generation': {
            'A': ClientTempDetails['A'],
        }
    }

    if computation_result_301:
        K_client = call_compute_a_api(f"{Endpoint}computeA/compute-K_client?password={password}", data)
        ClientTempDetails["K_client"] = K_client
        print("K ---> ", ClientTempDetails["K_client"])

    print("Completion of 301")

    # Step 4: Handle 401
    M_1 = call_compute_a_api(f"{Endpoint}computeA/compute-K_client/computeM1", data)
    print("M1: ---> ", M_1)

    computation_result_401 = await login_401(uri, M_1)
    print(computation_result_401)

    # Step 5: Final step
    M = call_compute_a_api(f"{Endpoint}computeA/compute-K_client/computeM?M_1={M_1}", data)
    print("M: ---> ", M)

    computation_result_501 = await login_501(uri, str(M))
    print(computation_result_501)


async def register():
    Endpoint = "http://localhost:2004/"

    name = input("Enter your name: ")
    password_input = input("Enter your password: ")
    if not is_int(password_input):
        password = password_input
    else:
        password = int(password_input)

    url = f"{Endpoint}signup"
    data = {"username": name, "password": password}
    headers = {"Content-Type": "application/json"}

    signup_response = requests.post(url, data=json.dumps(data), headers=headers)

    response_content = json.loads(signup_response.text)
    response_content = json.loads(response_content['message'])

    for key in response_content:
        ClientDetails[key] = response_content[key]
    print(ClientDetails)
    
    url = f"{Endpoint}upload"
    headers = {"Content-Type": "application/json"}
    response = requests.post(url, data=json.dumps(ClientDetails), headers=headers)

    # Check the response
    if response.status_code == 200:
        print("Signup Success")
    elif response.status_code == 409:
        print("Signup Failure")
    else:
        print("Unexpected response:", response.status_code)
        print(response.text)


if __name__ == "__main__":
    value = int(input("1. Register\n2. Login\n3. Exit\n"))
    if value == 2:
        asyncio.run(main())
    elif value == 1:
        asyncio.run(register())
    else:
        print("Exiting...")
