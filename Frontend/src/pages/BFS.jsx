import { useState } from 'react';
import { Box, Container, Input, Button, Flex, Center } from '@chakra-ui/react';
import bfsImg from '../assets/bfs.png';

const BFS = () => {
  const [start, setStart] = useState('');
  const [goal, setGoal] = useState('');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    setLoading(true);
    const response = await fetch(`http://localhost:8080/?src=${encodeURIComponent(start)}&dest=${encodeURIComponent(goal)}`);
    if (response.ok) {
      const data = await response.json();
      setResult(data);
    } else {
      alert('Failed to find path.');
    }
    setLoading(false);
  };

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
              <Input 
                value={start}
                onChange={(e) => setStart(e.target.value)}
                placeholder="Enter the start article" 
                ml={20} 
                mr={5} 
                borderColor="#214a6d" 
                borderWidth="2px" 
                width="300px" 
              />
            </Center>
            <Center>
              <Input 
                value={goal}
                onChange={(e) => setGoal(e.target.value)}
                placeholder="Enter the goal article" 
                borderColor="#214a6d" 
                borderWidth="2px" 
                width="300px" 
              />
            </Center>
          </Flex>
        </Container>

        <Container mt={10} mb={40} fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Button 
                onClick={handleSearch} 
                isLoading={loading}
                bgColor="#214a6d" 
                color="white" 
                ml={20}
              > 
                Start 
              </Button>
            </Center>
          </Flex>
        </Container>
      </Box>
    </div>
  )
}

export default BFS;