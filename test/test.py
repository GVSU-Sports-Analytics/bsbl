import requests


APIKEY = "test"
ENDPOINT = "http://127.0.0.1:8080/"


def test_index_route() -> None:
    response = requests.post(
            url=ENDPOINT,
            json={
                "api_key": APIKEY,
                "team_1_name": "Cincinnati Reds",
                "team_1_yr": "1975",
                "team_2_name": "Kansas City Royals",
                "team_2_yr": "2015",
            }
    )

    print(response.text)


if __name__ == "__main__":
    test_index_route()

