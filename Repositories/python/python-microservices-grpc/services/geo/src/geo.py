import math
import time
from concurrent import futures
from dataclasses import dataclass

import grpc
from data.data_load_module import load_data, build_hotel_geo_data, hotel_geo_data
from proto import geo_pb2, geo_pb2_grpc

GEO_SERVICE_ADDRESS = "[::]:8080"
EARTH_RADIUS_KM = 6371.0
MAX_SEARCH_RADIUS_KM = 10  # limit to 10 km


@dataclass
class Point:
    point_latitude: float
    point_longitude: float


def haversine_distance(coord1, coord2):
    lat1, lon1 = map(math.radians, [coord1.point_latitude, coord1.point_longitude])
    lat2, lon2 = map(math.radians, [coord2.point_latitude, coord2.point_longitude])

    dlat = lat2 - lat1
    dlon = lon2 - lon1

    a = math.sin(dlat / 2) * math.sin(dlat / 2) + math.cos(lat1) * math.cos(
        lat2
    ) * math.sin(dlon / 2) * math.sin(dlon / 2)
    c = 2 * math.atan2(math.sqrt(a), math.sqrt(1 - a))

    return EARTH_RADIUS_KM * c


def find_nearby_hotels(point, radius):
    nearby_hotels = []

    for hotel_id, hotel_info in hotel_geo_data.items():
        hotel_coord = Point(
            point_latitude=hotel_info["lat"], point_longitude=hotel_info["lon"]
        )
        distance = haversine_distance(point, hotel_coord)

        if distance <= radius:
            nearby_hotels.append(hotel_id)

    return nearby_hotels


class GeoServicer(geo_pb2_grpc.GeoServicer):
    def Nearby(self, request, context):
        point = Point(point_latitude=request.lat, point_longitude=request.lon)
        hotel_ids = find_nearby_hotels(point, MAX_SEARCH_RADIUS_KM)
        return geo_pb2.Result(hotelIds=hotel_ids)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    geo_pb2_grpc.add_GeoServicer_to_server(GeoServicer(), server)
    server.add_insecure_port(GEO_SERVICE_ADDRESS)
    load_data()
    build_hotel_geo_data()
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)
