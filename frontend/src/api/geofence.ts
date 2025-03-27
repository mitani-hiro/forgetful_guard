import { apiClient } from "./client";
import { components } from "./openapi.gen";

export const createGeofence = async (
  polygon: [number, number][][],
  token: string | null
) => {
  console.log("polygon: ", polygon);

  const geofence: components["schemas"]["Geofence"] = {
    title: "hoge title",
    userID: 999,
    polygon: polygon,
    deviceToken: token || "",
  };

  const response = await apiClient.post("/api/geofence", geofence);
};
