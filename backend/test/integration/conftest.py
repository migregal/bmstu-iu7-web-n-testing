import pytest as pt

from api.cube import Cube, CubeSupport

from config import *


@pt.fixture(scope="session", autouse=True)
def prefill_users():
    cube = CubeSupport(Cube(CUBE_HOST))
    users = [
        {
            "email": USER_EMAIL,
            "password": USER_PASSWORD,
            "username": USER_USERNAME,
            "fullname": USER_FULLNAME,
        }
    ]

    for u in users:
        cube.registration(u["email"], u["password"], u["username"], u["fullname"])
