
class Grafo:
    def __init__(self, es_dirigido=False, vertices=None):
        self.dirigido = es_dirigido
        self.vertices = {}
        if vertices:
            for vertice in vertices:
                self.vertices[vertice] = {}
    
    def agregar_vertice(self, v):
        if v in self.vertices:
            raise ValueError(f"Ya hay un vertice {v} en el grafo")
        self.vertices[v] = {}
    
    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError(f"No hay un vertice {v} en el grafo")
        self.vertices.pop(v)
        for adyacentes in self.vertices.values():
            for vertice in list(adyacentes.keys()):
                if vertice == v:
                    adyacentes.pop(v)
    
    def agregar_arista(self, v, w, peso=1):
        if v not in self.vertices or w not in self.vertices:
            raise ValueError(f"No hay vertice {v} o {w} en el grafo")
        if w in self.vertices[v]:
            raise ValueError(f"El vertice {v} ya tiene como adyacente al vertice {w}")
        
        self.vertices[v][w] = peso
        if not self.dirigido:
            self.vertices[w][v] = peso
    
    def estan_unidos(self, v, w):
        if v not in self.vertices or w not in self.vertices:
            return False
        if w in self.vertices[v]:
            return True
        if not self.dirigido and v in self.vertices[w]:
            return True
        return False
    
    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise ValueError(f"El vertice {v} no tiene como adyacente el vertice {w}")
        return self.vertices[v][w]

    def obtener_vertices(self):
        return list(self.vertices.keys())
    
    def vertice_aleatorio(self):
        vertices = self.obtener_vertices()
        return vertices[0] if vertices else None

    def adyacentes(self, v):
        return list(self.vertices[v].keys())
