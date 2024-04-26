import React, { useState } from 'react';
import { Box, Container, Input, Button, Flex, Center, useToast } from '@chakra-ui/react';
import idsImg from '../assets/ids.png';

const IDS = () => {
  const [start, setStart] = useState('');
  const [goal, setGoal] = useState('');
  const [result, setResult] = useState(null);
  const [executionTime, setExecutionTime] = useState(null);
  const [visitedCount, setVisitedCount] = useState(null);
  const [length, setLength] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const toast = useToast(); 

  const checkArticleExistence = async (title) => {
    const url = `https://en.wikipedia.org/w/api.php?action=query&format=json&titles=${encodeURIComponent(title)}&origin=*`;
    try {
        const response = await fetch(url);
        const data = await response.json();
        const page = data.query.pages;
        const pageId = Object.keys(page)[0];
        return pageId !== "-1";
    } catch (error) {
        console.error('Error checking article:', error);
        return false;
    }
  };

  const handleSearch = async () => {
    if (start.trim() === '' || goal.trim() === '') {
        toast({
            title: "Validation Error",
            description: "Both start and goal titles are required.",
            status: "error",
            duration: 9000,
            isClosable: true,
            position: "top"
        });
        return;
    }

    const startExists = await checkArticleExistence(start);
    const goalExists = await checkArticleExistence(goal);

    if (!startExists || !goalExists) {
        toast({
            title: "Article Not Found",
            description: `The article "${!startExists ? start : goal}" does not exist.`,
            status: "error",
            duration: 9000,
            isClosable: true,
            position: "top"
        });
        return;
    }

    setLoading(true);
    setError(null);
    try {
        const response = await fetch(`http://localhost:8081/?startTitle=${encodeURIComponent(start)}&goalTitle=${encodeURIComponent(goal)}`);
        const data = await response.json();
        if (!response.ok) {
            throw new Error(data.error || `Failed to fetch path. Status code: ${response.status}`);
        }
        if (!data.paths || data.paths.length === 0) {
            throw new Error("No path found between the articles.");
        }
        setResult(data.paths);
        setExecutionTime(data.timeTaken);
        setVisitedCount(data.visited);
        setLength(data.length);
    } catch (error) {
        console.error('Error:', error);
        setError(error.message);
        toast({
            title: "Error",
            description: error.message,
            status: "error",
            duration: 9000,
            isClosable: true,
            position: "top"
        });
    } finally {
        setLoading(false);
    }
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
              <Button onClick={handleSearch} isLoading={loading} bgColor="#465a3b" color="white"> Start </Button>
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