#!/usr/local/bin/python
# -*- coding: utf-8 -*-

import numpy as np
import sympy as sp

from sympy.abc import alpha, sigma, k, l, A, b


def Koef(EqS, d):
    """
  Функция на вход принимает упрощенное уравнение EqS с выделенными функциями при коэффицентах "сигма" и словарь d со значениями переменных в уравнении.
  На выходе получаем массив, содержащий конекретные значения коэффицентов полинома, индекс в массиве соответствует номеру коэффицента при степени "сигма".
    """
    s = np.array([])
    for i in range(0, len(EqS.keys())):
        rt = (sp.expand(EqS[sigma ** i]))
        s = np.append(s, rt.evalf(subs=d))
    return np.clongdouble(s)


def polyOdnorod():
    """
        Описывает случай для произвольной глубины. У полинома 4 ветви корней с учетом кратности
    """
    D, Dt, Del, Rm, mu, rho, L_2, O, b_0xy, lnH_0, skl, lk = sp.symbols(
        'D Dt Del Rm mu rho L_2 O  b_0xy lnH_0 skl lk')  # L_2 - двумерный оператор Лапласса, skl = s(k,l) = [b_0x*k +b_0y*l],  b_0xy = [b_0x + b_0y], lk=[k**2+l**2]

    # Gолный вид первого слагаемого
    I = (Dt * Dt + alpha * alpha)
    # Подставляем вместо операторов значения полученные в результате применения к кси самих операторов. Таким образом уравнение сводиться к алгебраическому
    u1 = I

    II = ((Dt - Del / Rm) * Dt - D * D / (mu * rho)) * Del
    # Подставляем вместо операторов значения полученные в результате применения к кси самих операторов. Таким образом уравнение сводиться к алгебраическому
    u2 = II

    Eq = D * u1 * u1 * u2

    sqf, Zx, Zy, Zz, b_0x, b_0y = sp.symbols('sqf Zx Zy Zz b_0x b_0y')
    Eq = Eq.subs([(D, 1j * (b_0x * k + b_0y * l)), (Del, (-1) * (k * k + l * l)),
                  (Dt, -1j * sigma)])
    return sp.collect(sp.expand(Eq), sigma, evaluate=False)


def polyNonOdnorod():
    """
        Функция упращает уравнение и возращает полином от сигма в виде словаря,
        где при кажлом сигма_i имеется функциональная зависимость. Она описывает
         случай переменной глубины.
    """
    D, Dt, Del, Rm, mu, rho, L_2, O, b_0xy, lnH_0, skl, lk = sp.symbols(
        'D Dt Del Rm mu rho L_2 O  b_0xy lnH_0 skl lk')  # L_2 - двумерный оператор Лапласса, skl = s(k,l) = [b_0x*k +b_0y*l],  b_0xy = [b_0x + b_0y], lk=[k**2+l**2]

    # Gолный вид первого слагаемого
    I = (Dt * Dt + alpha * alpha)
    # Подставляем вместо операторов значения полученные в результате применения к кси самих операторов. Таким образом уравнение сводиться к алгебраическому
    u1 = I

    II = ((Dt - Del / Rm) * Dt - D * D / (mu * rho)) * Del
    # Подставляем вместо операторов значения полученные в результате применения к кси самих операторов. Таким образом уравнение сводиться к алгебраическому
    u2 = II

    Btil, Htil, b_0x, b_0y = sp.symbols('Btil Htil b_0x b_0y')

    u3 = Btil / Htil * (mu * rho * mu * rho)
    Eq = D * u1 * u1 * u2 - u3

    Eq = Eq.subs([(D, 1j * (b_0x * k + b_0y * l)), (Del, (-1) * (k * k + l * l)),
                  (Dt, -1j * sigma), (Btil, 1), (Htil, 1)])
    return sp.collect(sp.expand(Eq), sigma, evaluate=False)


def diskr(mass, h=0.1):
    """
    Функция на вход получает массив из начала и конца отрезка спарамтером дискретизации.
    На выходе получается массив значений от конда до начала с шагом h
    """
    s = np.array([])
    for i in range(int((mass[-1] - mass[0]) / h) + 1):
        s = np.append(s, mass[0] + h * i)
    return s


class MathModels:

    def __init__(self, k, l, b_0x, b_0y, alpha, A, H_0, g, Rm, mu, rho, h=0.1):
        self.k = k
        self.l = l
        self.b_0x = b_0x
        self.b_0y = b_0y
        self.alpha = alpha
        self.A = A
        self.H_0 = H_0
        self.g = g
        self.Rm = Rm
        self.mu = mu
        self.rho = rho
        self.b = b
        self.h = h

        flac = 0
        param = ''
        value = 0
        for key in self.__dict__:
            value_temp = self.__dict__[key]
            if isinstance(value_temp, list):
                flac += 1
                param = key
                value = value_temp
        if flac == 1:
            self.__dict__[param] = diskr(value, self.h)
        else:
            if flac == 0:
                print('Ошибка! Подайте значение "X" в виде "[X]"')
            else:
                print("Ошибка! Подайте только один параметр в качестве отрезка")

        copy_dict = self.__dict__.copy()
        copy_dict.pop(param)

        global_list = []
        temp_dict = {}

        for value in self.__dict__[param]:
            temp_dict[param] = value
            temp_dict.update(copy_dict)
            global_list.append(temp_dict)
            temp_dict = {}

        self.global_list = global_list

    def __str__(self):
        return str(self.__dict__)

    def create_list_parametrs(self, equation_form):
        list_final = []
        gloal_list = self.global_list

        for dict in gloal_list:
            list_final.append(Koef(equation_form, dict))

        return list_final


def processingForMainGo(list):
    # Функция на вход получает numpy.array и преобразовывает его в строку.

    str1 = str(list)
    str1 = str1.replace('dtype=complex128)]', '').replace('[', '').replace('j', 'i').replace('  ',
                                                                                             '').replace('\n',
                                                                                                         '').replace(
        'array(', '').replace('],', '').replace(' ', '')
    strFinal = ''

    str_pre = str1.split('dtype=complex128),')

    for i in str_pre:
        if i != '':
            strFinal += i
            strFinal += ';'
    return strFinal+'end'


if __name__ == "__main__":
    # print(polyNonOdnorod())
    # print("="*1200)
    test = MathModels(1, 2, 3, [1, 2, 3], 5, 1, 2, 1, 1, 1, 1, h=1)

    list_for_go = test.create_list_parametrs(polyNonOdnorod())
    print(processingForMainGo(list_for_go))
    # print('ghb'+'b')

    # print(Koef(polyNonOdnorod(), test.global_list[0]))
