import { Box, Text, Container } from "@chakra-ui/react";
import { extendTheme, ChakraProvider, GlobalStyle } from "@chakra-ui/react";

import HelloGraduationSans from "../assets/HelloGraduationSans-d9enl.ttf";

const theme = extendTheme({
  fonts: {
    body: "HelloGraduationSans",
    heading: "HelloGraduationSans",
  },
});


export default function Dashboard() {
  return (
    <ChakraProvider theme={theme}>
      <GlobalStyle
        styles={{
          "@font-face": {
            fontFamily: "HelloGraduationSans",
            src: `url(${HelloGraduationSans}) format("truetype")`,
          },
        }}
      />
      <Box bgGradient="linear(to-b, white, pink.100, pink)" minHeight="100vh" pt="30px">
        <Container>
          <img src="../../public/img/title.png" alt="title"></img>
          <img src="../../public/img/ppg.png" alt="ppg"></img>
        </Container>
        <Text
          mr="20px"
          ml="20px"
          mx="450px"
          alignItems="center"
          textAlign="center"
          fontFamily="body"
        >
          WikiRace or Wiki Game is a game involving Wikipedia, a free online encyclopedia managed by various volunteers worldwide,
          where players start at a Wikipedia article and must navigate through other articles on Wikipedia (by clicking on links within each article)
          to reach another pre-determined article within the shortest time or with the fewest clicks (articles).
        </Text>
      </Box>
    </ChakraProvider>
  );
}
