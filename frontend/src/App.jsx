import { Box, Button, Heading, VStack } from "@chakra-ui/react";
import { Link } from "react-router-dom";

function App() {
  return (
    <VStack spacing={5} p={10}>
      <Heading>Welcome to Hospital Management</Heading>
      <Box>
        <Button as={Link} to="/login" colorScheme="teal" mr={3}>
          Login
        </Button>
        <Button as={Link} to="/register" colorScheme="blue">
          Register
        </Button>
      </Box>
    </VStack>
  );
}

export default App;
