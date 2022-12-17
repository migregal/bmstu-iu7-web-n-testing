import requests

import logging
from http.client import HTTPConnection

log = logging.getLogger("urllib3")
log.setLevel(logging.DEBUG)

ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)
log.addHandler(ch)

HTTPConnection.debuglevel = 1

import allure

from hamcrest import (
    assert_that,
    equal_to,
    not_,
    greater_than,
    greater_than_or_equal_to,
    less_than_or_equal_to,
    all_of,
)


class Cube:
    def __init__(self, host: str):
        self.host = host

    def registration(self, email: str, password: str, username: str, fullname: str):
        body = {
            "email": email,
            "password": password,
            "username": username,
            "fullname": fullname,
        }

        return requests.post(f"{self.host}/api/v1/registration", json=body)

    def login(self, email: str, password: str, session: requests.Session = None):
        session = session or requests.session()

        body = {
            "email": email,
            "password": password,
        }

        return session.post(f"{self.host}/api/v1/login", json=body)

    def users(
        self, page: int = 0, per_page: int = 10, session: requests.Session = None
    ):
        session = session or requests.session()

        params = {"page": page, "per_page": per_page}

        return session.get(f"{self.host}/api/v1/users", params=params)

    def models(
        self, page: int = 0, per_page: int = 10, session: requests.Session = None
    ):
        session = session or requests.session()

        params = {"page": page, "per_page": per_page}

        return session.get(f"{self.host}/api/v1/models", params=params)


class CubeSupport:
    def __init__(self, cube: Cube):
        self.cube = cube

    @allure.step("Регистрируемся, ожидаем успех")
    def registration(self, email: str, password: str, username: str, fullname: str):
        r = self.cube.registration(email, password, username, fullname)

        assert_that(r.status_code, equal_to(200))
        assert_that(r.json()["token"], not_(equal_to("")))

        return r

    @allure.step("Авторизуемся, ожидаем успех")
    def login(self, email: str, password: str, session: requests.Session = None):
        r = self.cube.login(email, password, session)

        assert_that(r.status_code, equal_to(200))
        assert_that(r.json()["token"], not_(equal_to("")))

        if session is not None:
            session.headers.update({"Authorization": f'Token {r.json()["token"]}'})

        return r

    @allure.step("Авторизуемся, ожидаем ошибку")
    def login_fail(self, email: str, password: str):
        r = self.cube.login(email, password)

        assert_that(r.status_code, equal_to(401))
        assert_that(r.json()["message"], equal_to("incorrect Username or Password"))

        return r

    @allure.step("Получаем список пользователей, ожидаем успех")
    def users(
        self, page: int = 0, per_page: int = 10, session: requests.Session = None
    ):
        r = self.cube.users(page, per_page, session=session)

        assert_that(r.status_code, equal_to(200))
        assert_that(
            len(r.json()["infos"]),
            all_of(greater_than(0), less_than_or_equal_to(per_page)),
        )

        return r

    @allure.step("Получаем список моделей, ожидаем успех")
    def models(
        self, page: int = 0, per_page: int = 10, session: requests.Session = None
    ):
        r = self.cube.users(page, per_page, session=session)

        assert_that(r.status_code, equal_to(200))
        assert_that(
            len(r.json()["infos"]),
            all_of(greater_than_or_equal_to(0), less_than_or_equal_to(per_page)),
        )

        return r
