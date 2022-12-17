import allure
import pytest as pt
from hamcrest import *

from api.cube import Cube, CubeSupport

from config import *


@pt.fixture(scope="module")
def cube():
    return CubeSupport(Cube(CUBE_HOST))


@allure.suite("api/v1/registration")
@allure.title("Если данные корректны, то должно зарегистрировать")
@pt.mark.parametrize(
    "email, password, username, fullname",
    [
        pt.param(
            NEW_USER_EMAIL,
            NEW_USER_PASSWORD,
            NEW_USER_USERNAME,
            NEW_USER_FULLNAME,
            id="default",
        )
    ],
)
def test_registration(cube: Cube, email, password, username, fullname):
    cube.registration(email, password, username, fullname)
