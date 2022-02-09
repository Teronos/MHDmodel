#!/usr/local/bin/python
# -*- coding: utf-8 -*-

import numpy as np
import json
import argparse
from lib.FunksForGraphic import PlZAv

def runParser():
  parser = argparse.ArgumentParser()
  parser.add_argument('-j','--json-string' )
  parser.add_argument('-p','--parametrs-for-create-graphics' )
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
    list_info_poly = json.loads(namespace.__dict__['json_string'])
    dataParametrs = json.loads(namespace.__dict__['parametrs_for_create_graphics'])
    
    PlZAv(loadingInfoPolyForPaint(list_info_poly), dataParametrs['-v'], dataParametrs['-H'][0])
    
    
   
    
   
    
    

        
