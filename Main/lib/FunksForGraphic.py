import numpy as np
import matplotlib.pyplot as plt

def diskr(mass, h=0.1):
    """
    Функция на вход получает массив из начала и конца отрезка спарамтером дискретизации.
    На выходе получается массив значений от конда до начала с шагом h
    """
    s = np.array([])
    for i in range(int((mass[-1] - mass[0]) / h) + 1):
        s = np.append(s, mass[0] + h * i)
    return s

def Razdel(mass):
    ar = np.array([])
    s = 0
    for i in range(0, len(mass) - 1):
        s = s + mass[i]
        ar = np.append(ar, s)
    return np.int_(ar)

def graf(mass, ar, h):
    arX = np.array([])
    arY = np.array([])
    ar = diskr(ar, h)
    for i in range(0, len(mass)):
        for j in mass[i]:
            arX = np.append(arX, j)
            arY = np.append(arY, ar[i])
    return arX, arY


def PlZAv(zav, mass, h):
    s = np.split(zav[0], Razdel(zav[1]))
    paintN(s, mass, h)
    paint3D(s, mass, h)


def paintN(mass, ar, h):
    arX, arY = graf(mass, ar, h)

    fig, (ax1, ax2) = plt.subplots(
        nrows=1, ncols=2,
        figsize=(20, 10)
    )

    ax1.set_title(r'$Re(\sigma) = \sigma_1$')
    ax1.set_xlabel('$x$')
    ax1.scatter(arY, arX.real)
    ax1.plot([ar[0], ar[1] + 1], [0, 0], color='black', linestyle='solid')

    ax2.set_title(r'$Im(\sigma) =\sigma_2$')
    ax2.set_xlabel('$x$')
    ax2.scatter(arY, arX.imag)
    ax2.plot([ar[0], ar[1] + 1], [0, 0], color='black', linestyle='solid')
    plt.show()


def paint3D(mass, ar, h):
    arX, arY = graf(mass, ar, h)

    x = arX.real
    y = arX.imag
    z = arY

    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')
    # ax.scatter(x, y, z)
    ax.scatter3D(x, y, z, c=z, cmap='RdBu_r')
    plt.show()
