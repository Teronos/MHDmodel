from os import name
import matplotlib.pyplot as plt
import numpy as np
import lib.Funks as lf


if __name__ == "__main__":
    m = np.array([[2, 3j, 5, -7], [1, 2, 3, 4j, 1j]])
    lf.PlZAv(m, [0, 2], 0.5)
