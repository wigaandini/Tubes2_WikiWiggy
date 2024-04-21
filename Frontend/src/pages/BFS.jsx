import { Box, Container, Input, Button, Flex, Center } from "@chakra-ui/react";
import bfsImg from "../assets/bfs.png";

const BFS = () => {
  return (
    <div>
      <Box bgGradient="linear(to-b, white, #cfe8fb, #8facc4)" minHeight="90vh" display="flex" flexDirection="column" alignItems="center" justifyContent="center">
        <Container mt={20}>
          <Flex justifyContent="center">
            <Center>
              <img 
                src={bfsImg} 
                alt="BFS" 
                style={{ maxWidth: '800px', height: 'auto' }}
              />
            </Center>
          </Flex>
        </Container>

        <Container mt="50px" fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Input placeholder="Enter the start article" ml={20} mr={5} borderColor="#214a6d" borderWidth="2px" width="300px" />
            </Center>
            <Center>
              <Input placeholder="Enter the goal article" borderColor="#214a6d" borderWidth="2px" width="300px" />
            </Center>
          </Flex>
        </Container>

        <Container mt={10} mb={40} fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Button bgColor="#214a6d" color="white" ml={20}> Start </Button>
            </Center>
          </Flex>
        </Container>
      </Box>
    </div>
  )
}

export default BFS;
