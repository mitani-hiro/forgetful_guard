import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  FlatList,
  ActivityIndicator,
  StyleSheet,
} from "react-native";
import { StackScreenProps } from "@react-navigation/stack";
import axios from "axios";
import { RootStackParamList } from "../../app";
import { useTaskStore } from "../../store/task";

type Props = StackScreenProps<RootStackParamList, "TaskList">;

const TaskListScreen = ({ navigation }: Props) => {
  const { tasks, fetchTasks } = useTaskStore();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchTasks();
    setLoading(false);
  }, []);

  if (loading) {
    return <ActivityIndicator size="large" style={styles.loader} />;
  }

  return (
    <View style={styles.container}>
      <FlatList
        data={tasks}
        keyExtractor={(item) => item.id.toString()}
        renderItem={({ item }) => (
          <View style={styles.taskItem}>
            <Text style={styles.title}>{item.title}</Text>
            <Text style={styles.description}>{item.description}</Text>
          </View>
        )}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: { flex: 1, padding: 16 },
  loader: { flex: 1, justifyContent: "center", alignItems: "center" },
  taskItem: {
    padding: 16,
    marginBottom: 8,
    backgroundColor: "#f9f9f9",
    borderRadius: 8,
  },
  title: { fontSize: 18, fontWeight: "bold" },
  description: { fontSize: 14, color: "#666" },
});

export default TaskListScreen;
