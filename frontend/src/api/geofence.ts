import { apiClient } from "./client";

export const createGeofence = async (
  polygon: [number, number][][],
  token: string | null
) => {
  console.log("polygon: ", polygon);

  const response = await apiClient.post(
    "api/geofence",
    JSON.stringify({
      title: "hoge title",
      userId: 999,
      polygon: polygon,
      deviceToken: token,
    })
  );

  console.log("api response: ", response);
};
