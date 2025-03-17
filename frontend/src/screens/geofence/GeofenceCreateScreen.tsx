import React from "react";
import MapboxGL from "@rnmapbox/maps";
import { View, Button } from "react-native";
import { StackScreenProps } from "@react-navigation/stack";
import { RootStackParamList } from "../../App";
import { useGeofenceStore } from "../../store/geofence";

type Props = StackScreenProps<RootStackParamList, "GeofenceCreate">;

MapboxGL.setAccessToken("YOUR_MAPBOX_ACCESS_TOKEN");

const GeofenceCreateScreen = ({ navigation }: Props) => {
  const { polygonCoordinates, addPoint, resetPolygon, registerGeofence } =
    useGeofenceStore();

  const handleMapPress = async (event: any) => {
    const coords = event.geometry.coordinates;
    console.log("handleMapPress coords: ", coords);

    if (polygonCoordinates.length >= 4) {
      console.log("最大4点まで選択可能です");
      return;
    }

    addPoint([coords[0], coords[1]]);
  };

  return (
    <View style={{ flex: 1 }}>
      <MapboxGL.MapView style={{ flex: 1 }} onPress={handleMapPress}>
        <MapboxGL.Camera
          zoomLevel={14}
          centerCoordinate={[139.6917, 35.6895]}
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
