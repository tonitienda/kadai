import { Typography, Container } from "@mui/material";
import TaskList from "../components/TaskList";

export default function Home() {
  return (
    <main>
      <Container maxWidth="sm">
        <TaskList />
      </Container>
    </main>
  );
}
