"""
Guido's Lasagna Cooking Module

This module provides functions to assist in calculating the cooking and preparation times for 
Guido's lasagna. It includes the expected baking time, preparation time per layer, and functions to 
determine the remaining baking time and total elapsed time.
"""

EXPECTED_BAKE_TIME: int = 40 # The total expected baking time for the lasagna in minutes.
PREPARATION_TIME: int = 2 # The time required to prepare each layer of the lasagna in minutes.

def bake_time_remaining(elapsed_bake_time: int) -> int:
    """
    Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes import the actual minutes the lasagna has been in the oven as an argument 
    and returns how many minutes the lasagna still needs to bake based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elapsed_bake_time

def preparation_time_in_minutes(number_of_layers: int) -> int:
    """
    Calculate the preparation time based on the number of layers.

    :param number_of_layers: int - number of layers in the lasagna.
    :return: int - total preparation time in minutes.
    """

    return PREPARATION_TIME * number_of_layers

def elapsed_time_in_minutes(number_of_layers: int, elapsed_bake_time: int) -> int:
    """
    Calculate the total elapsed time including preparation and baking.

    :param number_of_layers: int - number of layers in the lasagna.
    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - total elapsed time in minutes.
    """

    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
