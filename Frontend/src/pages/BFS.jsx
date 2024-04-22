import { useState } from 'react';
import { Box, Container, Input, Button, Flex, Center } from '@chakra-ui/react';
import bfsImg from '../assets/bfs.png';

const BFS = () => {
  const [start, setStart] = useState('');
  const [goal, setGoal] = useState('');
  const [result, setResult] = useState(null);
  const [executionTime, setExecutionTime] = useState(null);
  const [visitedCount, setVisitedCount] = useState(null);
  const [length, setLength] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null); // State for error handling

  const handleSearch = async () => {
    setLoading(true);
    setError(null); // Reset error state before making a new request
    try {
      const response = await fetch(`http://localhost:8080/?startTitle=${encodeURIComponent(start)}&goalTitle=${encodeURIComponent(goal)}`);
      if (response.ok) {
        const data = await response.json();
        if (data.timeTaken) {
          setResult(data.paths);
          setExecutionTime(data.timeTaken);
          setVisitedCount(data.visited);
          setLength(data.length);
        } else {
          throw new Error('Invalid response format: Attribute is missing');
        }
      } else {
        throw new Error('Failed to fetch path. Status code: ' + response.status);
      }
    } catch (error) {
      console.error('Error:', error); // Log the error to the console for debugging
      setError('Failed to fetch data. Please try again.'); // Set error message for user display
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
  
        {result && (
          <Container mt={5} fontFamily="monospace">
            <Flex direction="column" align="center">
              <Box mb={2}>
                <b>Path:</b> {result.join(' ➡️ ')}
              </Box>
              <Box mb={2}>
                <b>Time Taken:</b> {executionTime} ms
              </Box>
              <Box mb={2}>
                <b>Visited:</b> {visitedCount}
              </Box>
              <Box mb={2}>
                <b>Length:</b> {length}
              </Box>
            </Flex>
          </Container>
        )}
      </Box>
    </div>
  )
}

export default BFS;