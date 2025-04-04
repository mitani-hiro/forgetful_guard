import { create } from "zustand";
import { createGeofence } from "../api/geofence";

interface GeofenceState {
  polygonCoordinates: [number, number][]; // ポリゴンの頂点座標
  addPoint: (coord: [number, number]) => void;
  resetPolygon: () => void;
  createGeofence: (token: string | null) => Promise<void>;
}

export const useGeofenceStore = create<GeofenceState>((set, get) => ({
  polygonCoordinates: [],

  addPoint: (coord) =>
    set((state) => ({
      polygonCoordinates: [...state.polygonCoordinates, coord],
    })),

  resetPolygon: () => set({ polygonCoordinates: [] }),

  createGeofence: async (token) => {
    const { polygonCoordinates } = get();
    if (polygonCoordinates.length < 3) {
      alert("ジオフェンスには3点以上の頂点が必要です");
      return;
    }

    if (!token) {
      alert("デバイストークンが不正です");
      return;
    }

    try {
      const polygon = [[...polygonCoordinates, polygonCoordinates[0]]];

      await createGeofence(polygon, token);
      alert("ジオフェンスが登録されました");
    } catch (error) {
      alert(error);
    }
  },
}));
