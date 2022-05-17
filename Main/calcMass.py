#!/usr/local/bin/python
# -*- coding: utf-8 -*-

import numpy as np
import sympy as sp
import json as js
import argparse

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

def homogeneous():
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

def inhomogeneous():
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

def BegData(kl, b0xy, al_, A_, H_, g_, mu_, ro_, b_, Rm_, func):
    b_0x, b_0y, g, H_0, lnH_0, Rm, mu, ro = sp.symbols(
        'b_0x b_0y g H_0 lnH_0 Rm mu rho')
    h = H_
    logh = np.log(h)

    kl = np.array(kl)
    k_ = kl[0]
    l_ = kl[1]

    b0xy = np.array(b0xy)
    b0x = b0xy[0]
    b0y = b0xy[1]

    d = {k: k_, l: l_, alpha: al_, A: A_, H_0: h, g: g_, lnH_0: logh, Rm: Rm_,
         mu: mu_, ro: ro_, b_0x: b0x, b_0y: b0y, b: b_}
    jk = Koef(func, d)  # Постановка значения в скобки поли меняет тип исслудуемого уравнения
    return jk

def pars_json(str_input):
  dict_out = js.loads(str_input)
  for key in dict_out.keys():
    value = dict_out[key]
    if len(value) == 1:
      dict_out[key] = value[0]
  return dict_out

def diskr(mass, h=0.1):
    """
    Функция на вход получает массив из начала и конца отрезка спарамтером дискретизации.
    На выходе получается массив значений от конда до начала с шагом h
    """
    s = np.array([])
    for i in range(int((mass[-1] - mass[0]) / h) + 1):
        s = np.append(s, mass[0] + h * i)
    return s      

def global_list(dict_input):
  for key in dict_input:
    value_dict = dict_input[key]
    if isinstance(value_dict, list):
      param = key
  copy_dict = dict_input.copy()
  copy_dict.pop(param)
  global_list = []
  temp_dict = {}
  for value in diskr(dict_input[param], dict_input['h']):
    temp_dict[param] = value
    temp_dict.update(copy_dict)
    global_list.append(temp_dict)
    temp_dict = {}
  return global_list

def create_list_parametrs(list_dict, equation_form):
  list_final = []

  for dictt in list_dict:
    list_final.append(Koef(EqS=equation_form,d=dictt))

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
            strFinal += ";"
    return strFinal + "end"

def runParser():
  parser = argparse.ArgumentParser()
  parser.add_argument('-im','--inhomogeneous' )
  parser.add_argument('-hm','--homogeneous' )
  return parser

if __name__ == "__main__":
  parser = runParser()
  namespace = parser.parse_args()
  if namespace.inhomogeneous :
    list_for_go = create_list_parametrs(global_list(pars_json(namespace.inhomogeneous)), inhomogeneous())
    print(processingForMainGo(list_for_go))

  if namespace.homogeneous :
    list_for_go = create_list_parametrs(global_list(pars_json(namespace.homogeneous)), homogeneous())
    print(processingForMainGo(list_for_go))
