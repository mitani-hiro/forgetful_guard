import React, { useEffect } from "react";
import MapboxGL from "@rnmapbox/maps";
import { View, Button } from "react-native";
import { StackScreenProps } from "@react-navigation/stack";
import Geolocation from "react-native-geolocation-service";
import { RootStackParamList } from "../../App";
import { useGeofenceStore } from "../../store/geofence";
import { useTrackerStore } from "../../store/tracker";

type Props = StackScreenProps<RootStackParamList, "GeofenceCreate">;

MapboxGL.setAccessToken("YOUR_MAPBOX_ACCESS_TOKEN");

const GeofenceCreateScreen = ({ navigation }: Props) => {
  const { polygonCoordinates, addPoint, resetPolygon, registerGeofence } =
    useGeofenceStore();

  const { sendTracker } = useTrackerStore();

  useEffect(() => {
    const stopTracking = startTracking();
    return () => stopTracking(); // コンポーネントがアンマウントされたら停止
  }, []);

  const handleMapPress = async (event: any) => {
    const coords = event.geometry.coordinates;
    console.log("handleMapPress coords: ", coords);

    if (polygonCoordinates.length >= 4) {
      console.log("最大4点まで選択可能です");
      return;
    }

    addPoint([coords[0], coords[1]]);
  };

  const startTracking = () => {
    const interval = setInterval(sendCurrentPosition, 5000);
    return () => clearInterval(interval);
  };

  const sendCurrentPosition = async () => {
    const position = await getCurrentPosition();
    sendTracker([position.longitude, position.latitude]);
  };

  const getCurrentPosition = () => {
    return new Promise<{ latitude: number; longitude: number }>(
      (resolve, reject) => {
        Geolocation.getCurrentPosition(
          (position) => {
            resolve({
              latitude: position.coords.latitude,
              longitude: position.coords.longitude,
            });
          },
          (error) => reject(error),
          { enableHighAccuracy: true, timeout: 15000, maximumAge: 10000 }
        );
      }
    );
  };

  return (
    <View style={{ flex: 1 }}>
      <MapboxGL.MapView style={{ flex: 1 }} onPress={handleMapPress}>
        <MapboxGL.Camera
          zoomLevel={14}
          //centerCoordinate={[139.6917, 35.6895]}
          followUserLocation
        />

        <MapboxGL.UserLocation />

        {polygonCoordinates.length > 2 && (
          <MapboxGL.ShapeSource
            id="polygonSource"
            shape={{
              type: "Polygon",
              coordinates: [
                [...polygonCoordinates, polygonCoordinates[0]], // 最初と最後をつなぐ
              ],
            }}
          >
            <MapboxGL.FillLayer
              id="polygonFill"
              style={{ fillColor: "rgba(0, 0, 255, 0.3)" }}
            ></MapboxGL.FillLayer>
          </MapboxGL.ShapeSource>
        )}
      </MapboxGL.MapView>

      <View style={{ position: "absolute", bottom: 50, left: 20, right: 20 }}>
        <Button title="ジオフェンス登録" onPress={registerGeofence}></Button>
        <Button title="リセット" onPress={resetPolygon} color="red"></Button>
      </View>
    </View>
  );
};

export default GeofenceCreateScreen;
