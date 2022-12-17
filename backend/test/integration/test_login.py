import allure
import pytest as pt

from api.cube import Cube, CubeSupport

from config import *


@pt.fixture(scope="module")
def cube():
    return CubeSupport(Cube(CUBE_HOST))


@allure.suite("api/v1/login")
@allure.title("Если передать креды существующего пользователя, то должно залогинить")
@pt.mark.parametrize(
    "email, password", [pt.param(USER_EMAIL, USER_PASSWORD, id="default")]
)
def test_login(cube: Cube, email, password):
    cube.login(email, password)


@allure.suite("api/v1/login")
@allure.title("Если передать некорректные креды, то должно отдать ошибку")
@pt.mark.parametrize(
    "email, password",
    [pt.param(USER_EMAIL, USER_INCORRECT_PASSWORD, id="default")],
)
def test_login_fail(cube: Cube, email, password):
    cube.login_fail(email, password)
