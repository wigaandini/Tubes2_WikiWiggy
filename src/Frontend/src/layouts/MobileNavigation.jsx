import { MdClose } from 'react-icons/md';
import { Link} from 'react-router-dom';
import Logo from '../assets/blossom.png';
import { Box, Image, Fade, Text} from '@chakra-ui/react';

const MenuItem = ({ children, isLast, to = "/", ...rest }) => {
    return (
      <Link to={to} >
        <Box w={"50vw"} display="flex" alignItems={"center"} justifyContent={"center"} 
                  {...rest} >
          <Text fontSize="1.2em" fontWeight={"bold"} py={2} textColor={"white"}>
              {children}
          </Text>
        </Box>
      </Link>
    );
  };

const MobileNavigation = ({ open, closeMenu }) => {

    return (
        <Fade in={open} 
        transitionDuration={"0.5s"}
        transitionTimingFunction={'ease-in-out'}>

            {/* black layer */}
            { open ? 
            <Box onClick={() => closeMenu()}
            position="fixed" top = {0} left ={0} opacity={open? '0.5' : '0'} bg={"black"} 
            w="100%" h="100%" zIndex={1}/> 
            :
             <Box/>
            }

            <Box
            position = "fixed"
            top = {0}
            left =  {open? 0 : -300}
            py='15px'
            w = "300px"
            h = "100%"
            color='white'
            bg='#9EB6BF'
            shadow='md'
            zIndex ={15}
            transitionDuration={"0.5s"}
            transitionTimingFunction={'ease-in-out'}
            >
                <Box position = "relative" display={"flex"} flexDirection={"column"}
                justifyContent={"center"} alignItems={"center"}>
                    <Box position = "absolute" top = {2} right = {2} >
                        <MdClose size = {24} onClick={() => closeMenu()}/>
                    </Box>
                    <Image src={Logo} h={{base:"20vh"}} w= "auto" my="2em"/>

                    <Box bgColor={"#AEC6CF"} w={"full"} display={"flex"} flexDir={"column"} 
                    justifyContent={"center"} alignItems={"center"} py={4}>
                    <MenuItem to="" > Home </MenuItem>
                    <MenuItem to="profile" > About Us </MenuItem>
                    <MenuItem to="how-to-use"> How To Use </MenuItem>
                    </Box>
                </Box>
            </Box> 
        </Fade>
    );
};

export default MobileNavigation;