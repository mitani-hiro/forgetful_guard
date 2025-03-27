import React from "react";
import { registerRootComponent } from "expo";
import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";
import HomeScreen from "./screens/HomeScreen";
import TaskListScreen from "./screens/task/TaskListScreen";
import GeofenceCreateScreen from "./screens/geofence/GeofenceCreateScreen";

export type RootStackParamList = {
  Home: undefined;
  TaskList: undefined;
  GeofenceCreate: undefined;
};

const Stack = createStackNavigator<RootStackParamList>();

const App = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator id={undefined}>
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="TaskList" component={TaskListScreen} />
        <Stack.Screen name="GeofenceCreate" component={GeofenceCreateScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

registerRootComponent(App);
