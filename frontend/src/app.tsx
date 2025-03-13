import React from "react";
import { registerRootComponent } from "expo";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import TaskListScreen from "./screens/task/TaskListScreen";

const Stack = createNativeStackNavigator();

const App = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen
          name="TaskList"
          component={TaskListScreen}
          options={{ title: "タスク一覧" }}
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

registerRootComponent(App);
