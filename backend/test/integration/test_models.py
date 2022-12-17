import requests

import allure
import pytest as pt

from api.cube import Cube, CubeSupport

from config import *


@pt.fixture(scope="module")
def cube():
    return CubeSupport(Cube(CUBE_HOST))


@allure.suite("api/v1/users")
@allure.title("Если передать page и per_page, должно вернуть модели")
@pt.mark.parametrize(
    "email, password", [pt.param(USER_EMAIL, USER_PASSWORD, id="default")]
)
@pt.mark.parametrize("page, per_page", [pt.param(0, 10, id="default")])
def test_models(cube: Cube, email, password, page, per_page):
    session = requests.session()
    cube.login(email, password, session=session)
    cube.models(page, per_page, session=session)
