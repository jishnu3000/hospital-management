import { useNavigate } from "react-router-dom";
import { Box, Button, Text, VStack } from "@chakra-ui/react";

function Dashboard() {
  const navigate = useNavigate();

  const logout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  return (
    <VStack spacing={5} p={10}>
      <Text fontSize="2xl">Welcome to the Dashboard</Text>
      <Button colorScheme="red" onClick={logout}>
        Logout
      </Button>
    </VStack>
  );
}

export default Dashboard;
