import { useState } from 'react';
import { Box, Container, Input, Button, Flex, Center } from "@chakra-ui/react";
import idsImg from "../assets/ids.png";

const IDS = () => {
  const [start, setStart] = useState('');
  const [goal, setGoal] = useState('');
  const [result, setResult] = useState(null);
  const [executionTime, setExecutionTime] = useState(null);
  const [visitedCount, setVisitedCount] = useState(null);
  const [length, setLength] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleSearch = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`http://localhost:8081/?startTitle=${encodeURIComponent(start)}&goalTitle=${encodeURIComponent(goal)}`);
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
      console.error('Error:', error);
      setError('Failed to fetch data. Please try again.');
    }
    setLoading(false);
  };

  const getWikipediaLink = (title) => `https://en.wikipedia.org/wiki/${encodeURIComponent(title)}`;


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
              <Input 
                value={start} 
                onChange={(e) => setStart(e.target.value)} 
                placeholder="Enter the start article" 
                mr={8} 
                borderColor="#465a3b" 
                borderWidth="2px" 
                width="300px" 
              />
            </Center>
            <Center>
              <Input 
                value={goal} 
                onChange={(e) => setGoal(e.target.value)} 
                placeholder="Enter the goal article" 
                // mr={20} 
                borderColor="#465a3b" 
                borderWidth="2px" 
                width="300px" 
              />
            </Center>
          </Flex>
        </Container>

        <Container mt={10} mb={10} fontFamily="monospace">
          <Flex justifyContent="center">
            <Center>
              <Button onClick={handleSearch} isLoading={loading} bgColor="#465a3b" color="white" mr={20}> Start </Button>
            </Center>
          </Flex>
        </Container>
        
        {result && (
          <Container mt={5} fontFamily="monospace" fontSize={20}>
            <Flex direction="column" align="center">
              <Box mb={2}>
                Found path with length <b> {length} </b> from {' '} <b>
                <a href={getWikipediaLink(start)} style={{ textDecoration: 'underline', color: 'inherit' }}>
                  <span onMouseOver={(e) => { e.target.style.color = 'white'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = '#76856f'; }} 
                    onMouseOut={(e) => { e.target.style.color = 'inherit'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = 'transparent'; }}
                  >
                    {start}
                  </span>
                </a> {' '} </b>
                to {' '} <b>
                <a href={getWikipediaLink(goal)} style={{ textDecoration: 'underline', color: 'inherit' }}>
                  <span onMouseOver={(e) => { e.target.style.color = 'white'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = '#76856f'; }} 
                    onMouseOut={(e) => { e.target.style.color = 'inherit'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = 'transparent'; }}
                  >
                    {goal}
                  </span>
                </a> </b>
                using IDS : 
              </Box>
              <Box mb={2}>
                <b>
                  {result.map((article, index) => (
                    <span key={index}>
                      <a href={getWikipediaLink(article)} style={{ textDecoration: 'none', color: 'inherit' }}>
                      <span onMouseOver={(e) => { e.target.style.color = 'white'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = '#76856f'; }}
                        onMouseOut={(e) => { e.target.style.color = 'inherit'; e.target.style.textDecoration = 'underline'; e.target.style.backgroundColor = 'transparent'; }}
                      >
                          {article}
                        </span>
                      </a>
                      {index !== result.length - 1 && ' --> '}
                    </span>
                  ))}
                </b>
              </Box>
              <Box mb={2}>
                in <b> {executionTime} ms </b>
              </Box>
              <Box mb={10}>
                With total <b> {visitedCount}  </b> articles visited.
              </Box>
            </Flex>
          </Container>
        )}
      </Box>
    </div>
  )
}

export default IDS;