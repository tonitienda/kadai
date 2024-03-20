import { Typography, Container, Fab, Button } from "@mui/material";
import TaskList from "../components/TaskList";

export default function Home() {
  return (
    <main>
      <Container maxWidth="sm">
        <Typography variant="h1">Tasks</Typography>
        <Typography variant="h3">Manage your tasks with Kadai!</Typography>
        <Typography variant="body1">Your tasks</Typography>
        <TaskList />
      </Container>
    </main>
  );
}
