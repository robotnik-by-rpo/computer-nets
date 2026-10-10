import networkx as nx
from collections import Counter
from pathlib import Path
import statistics
import matplotlib.pyplot as plt

def load(edges_file: Path):
    G = nx.read_edgelist(edges_file, nodetype=int, create_using=nx.Graph())
    return G
def all_short_dist(G):
    return dict(nx.all_pairs_shortest_path_length(G))

def dist_distribution(dists: dict):
    counter = Counter()
    nodes = list(dists.keys())

    for i, u in enumerate(nodes):
        for v in nodes[i+1:]:
            d = dists[u].get(v)
            if d is not None and d > 0:
                counter[d] += 1

    return counter

def analyze_net(edges_file: Path):
    G = load(edges_file)
    n = G.number_of_nodes()
    m = G.number_of_edges()

    if not nx.is_connected(G):
        largest_con_com = max(nx.connected_components(G), key=len)
        G = G.subgraph(largest_con_com).copy()

    dist = all_short_dist(G)
    counter = dist_distribution(dist)

    return G, counter

def get_stats(counter):
    if not counter:
        return None
    dists = sorted(counter)

    total = sum(counter.values())

    mode = max(counter, key=counter.get)

    c = 0
    median = None
    for d in dists:
        c += counter[d]
        if c >= total / 2:
            median = d
            break

    mean = sum(d * c for d, c in counter.items()) / total
    return {
        "total_pairs": total,
        "mean": mean,
        "mode": mode,
        "median": median,
        "distribution": dict(sorted(counter.items())),
    }

def plot_distribution(counter, ego_id):
    dists = sorted(counter)
    counts = [counter[d] for d in dists]
    
    plt.figure(figsize=(10, 6))
    plt.bar(dists, counts, color="red")
    plt.xlabel("Длина кратчайшего пути")
    plt.ylabel("Количество пар вершин")
    plt.title(f"Распределение длин кратчайших путей (сеть {ego_id})")
    plt.xticks(dists)
    plt.grid(axis="y", alpha=0.3)
    plt.show()

if __name__ == "__main__":
    base = Path("./social-networks")
    ego_ids = [0, 107, 348, 414, 686, 698, 1684, 1912, 3437, 3980]

    summary = []
    for ego in ego_ids:
        edges = base / f"{ego}.edges"
        if not edges.exists():
            continue
        
        G, counter = analyze_net(edges)
        stats = get_stats(counter)
        if stats is None:
            continue

        print(f"\n=== Эго-сеть {ego} ===")
        print(f"Узлов: {G.number_of_nodes()}, рёбер: {G.number_of_edges()}")
        print(f"Пар: {stats['total_pairs']}")
        print(f"Среднее: {stats['mean']:.3f}")
        print(f"Мода: {stats['mode']}")
        print(f"Медиана: {stats['median']}")
        print("Распределение:", stats["distribution"])
        
        plot_distribution(counter, ego)
        summary.append((ego, stats))