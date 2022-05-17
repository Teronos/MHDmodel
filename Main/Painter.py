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

def argegate_to_complex_list(list_info_poly):
        listComplex = np.array([])
        list_amount_root = np.array([])
        for element in list_info_poly:
         
            list_amount_root = np.append(list_amount_root, len(element['MultiplicityRoots']))
            for value in element['MultiplicityRoots']:
                listComplex = np.append(listComplex, np.clongdouble(value.replace('i','j')))
         
        return (listComplex, list_amount_root) 

if __name__ == "__main__":
    parser = runParser()
    namespace = parser.parse_args()
    list_info_poly = json.loads(namespace.__dict__['json_string'])
    dataParametrs = json.loads(namespace.__dict__['parametrs_for_create_graphics'])
    m = argegate_to_complex_list(list_info_poly=list_info_poly)    
    PlZAv(m, dataParametrs['-v'], dataParametrs['-H'][0]) 
    
    
   
    
   
    
    

        
