#!/usr/bin/env python3

import sys
import grafo
import biblioteca
import funciones as f
from collections import deque

def main():
    argumentos = sys.argv[1:]
    if len(argumentos) != 2:
        raise Exception("Parametros invalidos")
    aeropuertos, vuelos = argumentos
    pila = deque()

    grafo_dinero = grafo.Grafo()
    grafo_tiempos = grafo.Grafo()
    grafo_frecuencias = grafo.Grafo()
    ciudades = {}  
    coordenadas = {}

    f.ingresar_aeropuertos(aeropuertos, grafo_dinero, grafo_tiempos, grafo_frecuencias, ciudades, coordenadas)
    f.ingresar_vuelos(vuelos, grafo_dinero, grafo_tiempos, grafo_frecuencias)

    for linea in sys.stdin:
        comando, parametros = linea.split(" ", 1)
        parametros = parametros.rstrip().split(",")

        if comando == "camino_mas":
            if len(parametros) != 3:
                print("Error de parametros")
                continue
            criterio, origen, destino = parametros
            if criterio == "barato":
                g = grafo_dinero
            elif criterio == "rapido":
                g = grafo_tiempos
            else:
                print("Error de parametros")
                continue
            if origen not in ciudades or destino not in ciudades:
                print("Error de parametros")
                continue
            res = f.camino_minimo(origen, destino, ciudades, g)
            pila.append(res)
            print(" -> ".join(res))

        elif comando == "camino_escalas":
            if len(parametros) != 2:
                print("Error de parametros")
                continue
            origen, destino = parametros
            if origen not in ciudades or destino not in ciudades:
                print("Error de parametros")
                continue
            res = f.camino_minimo_escalas(origen, destino, ciudades, grafo_dinero)
            pila.append(res)
            print(" -> ".join(res))

        elif comando == "centralidad":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            n = parametros[0]
            if not n.isdigit():
                print("Error de parametros")
                continue
            n = int(n)
            res = f.obtener_centralidad(grafo_frecuencias, n)
            print(", ".join(res))

        elif comando == "itinerario":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            ruta = parametros[0]
            grafo_ciudades = f.crear_itinerario(ruta)
            orden = biblioteca.orden_topologico_dfs(grafo_ciudades)
            print(", ".join(orden))
            for i in range(len(orden) - 1):
                origen, destino = orden[i], orden[i+1]
                res = f.camino_minimo_escalas(origen, destino, ciudades, grafo_frecuencias)
                print(" -> ".join(res))

        elif comando == "nueva_aerolinea":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            ruta = parametros[0]
            arbol = biblioteca.mst_prim(grafo_dinero)
            f.crear_rutas(ruta, arbol, grafo_tiempos, grafo_frecuencias)
            print("OK")

        elif comando == "exportar_kml":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            ruta = parametros[0]
            if len(pila) == 0:
                print("No se ha ejecutado un comando anteriormente")
                continue
            camino = pila.pop()
            f.crear_kml(ruta, camino, coordenadas)
            print("OK")

main()
