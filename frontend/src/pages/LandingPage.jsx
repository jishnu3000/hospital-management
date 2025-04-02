import { Box, Button, Heading, VStack } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

function LandingPage() {
  const navigate = useNavigate();

  return (
    <Box textAlign="center" mt="20">
      <VStack spacing={4}>
        <Heading>Hospital Management System</Heading>
        <Button colorScheme="teal" onClick={() => navigate("/login")}>
          Login
        </Button>
        <Button colorScheme="blue" onClick={() => navigate("/register")}>
          Register
        </Button>
      </VStack>
    </Box>
  );
}

export default LandingPage;
