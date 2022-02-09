#!/usr/local/bin/python
# -*- coding: utf-8 -*-

#import matplotlib.pyplot as plt
import numpy as np
import json
import argparse
#import lib.FunksForGraphic as lb

from lib.FunksForGraphic import PlZAv
""" 
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
"""

def runParser():
  parser = argparse.ArgumentParser()
  parser.add_argument('-j','--json-string' )
  parser.add_argument('-p','--parametrs-for-create-graphics' )
  #parser.add_argument('-H','--parametr_diskr' )
  return parser

def loadingInfoPolyForPaint(list_info_poly):
    listComplex = np.array([])
    listComplexAmountRoots = np.array([])
    for element in list_info_poly:
    
        for num in element['MultiplicityRoots']:
            listComplex = np.append(listComplex, np.clongdouble(num.replace('i','j')))
        listComplexAmountRoots = np.append(listComplexAmountRoots, len(element['MultiplicityRoots']))
    
    return (listComplex, listComplexAmountRoots)  

if __name__ == "__main__":
   

    parser = runParser()
    namespace = parser.parse_args()
    #print(namespace)
    list_info_poly = json.loads(namespace.__dict__['json_string'])
    #print(list_info_poly)
    #print(namespace.__dict__['parametrs_for_create_graphics'])
    dataParametrs = json.loads(namespace.__dict__['parametrs_for_create_graphics'])
    #print(dataParametrs['-H'][0])
    #print(dataParametrs['-v'])
    #print(dataParametrs["-H"])
    #print(loadingInfoPolyForPaint(list_info_poly))
    PlZAv(loadingInfoPolyForPaint(list_info_poly), dataParametrs['-v'], dataParametrs['-H'][0])
    #lb.PlZAv((np.array([1,2j,1]), np.array([1,1,1])), [1,2], 0.5)

    
   
    
   
    
    

        
