import { Box, Container, Input, Button, Flex, Center } from "@chakra-ui/react";
import idsImg from "../assets/ids.png";

const IDS = () => {
  return (
    <div>
      <Box bgGradient="linear(to-b, white, #d0e8c5, #a2b499)" minHeight="90vh" display="flex" flexDirection="column" alignItems="center" justifyContent="center">
        <Container mt={20}>
          <Flex justifyContent="center">
            <Center>
              <img 
                src={idsImg} 
                alt="IDS" 
                style={{ maxWidth: '800px', height: 'auto' }}
              />
            </Center>
          </Flex>
        </Container>

        <Container mt="50px" fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Input placeholder="Enter the start article" mr={5} borderColor="#465a3b" borderWidth="2px" width="300px" />
            </Center>
            <Center>
              <Input placeholder="Enter the goal article" mr={20} borderColor="#465a3b" borderWidth="2px" width="300px" />
            </Center>
          </Flex>
        </Container>

        <Container mt={10} mb={40} fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Button bgColor="#465a3b" color="white" mr={20}> Start </Button>
            </Center>
          </Flex>
        </Container>
      </Box>
    </div>
  )
}

export default IDS;