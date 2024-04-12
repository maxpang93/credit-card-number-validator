import requests

BASE_URL = "http://0.0.0.0:8081"


def ping():
    try:
        res = requests.get(f"{BASE_URL}/ping")
        res.raise_for_status()
        print(res.status_code)
        print(res.json())
    except Exception as e:
        print(e)


def validate_credit_card_number(num):
    data = {"creditCardNumber": num}
    try:
        res = requests.post(f"{BASE_URL}/validate", json=data)
        res.raise_for_status()
        print(res.status_code)
        print(res.json())
    except Exception as e:
        print(e)


if __name__ == "__main__":
    ping()
    validate_credit_card_number("1234-3456-5678")
    validate_credit_card_number("1234-3456_5678")
    validate_credit_card_number("123434565678")
    validate_credit_card_number("1234A34565678")
    validate_credit_card_number("abcdaadfawer")

    # valid card numbers
    validate_credit_card_number("17893729974")  # wikipedia example
    validate_credit_card_number("4283321000018291")  # maybank sample
