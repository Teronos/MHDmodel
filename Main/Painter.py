#!/usr/local/bin/python
# -*- coding: utf-8 -*-

import matplotlib.pyplot as plt
import numpy as np
import sys
import json

from lib.FunksForGraphic import PlZAv

def loadingInfoPolyForPaint():
    listComplex = np.array([])
    listComplexAmountRoots = np.array([])


    with open('./json_files/list_info_polynom.txt','r') as json_file:
            list_info_poly = json.load(json_file)    
    
    for element in list_info_poly:
    
        for num in element['MultiplicityRoots']:
            listComplex = np.append(listComplex, np.clongdouble(num.replace('i','j')))
        listComplexAmountRoots = np.append(listComplexAmountRoots, len(element['MultiplicityRoots']))
    
    return (listComplex, listComplexAmountRoots)

if __name__ == "__main__":
    with open('./json_files/variable_parametr_model.txt','r') as json_file:
        dataParametrs = json.load(json_file)

    PlZAv(loadingInfoPolyForPaint(), dataParametrs['variable'], dataParametrs['h'])

    
   
    
   
    
    

        
