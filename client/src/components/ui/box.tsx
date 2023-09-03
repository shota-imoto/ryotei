import { BoxProps, Box as MuiBox } from '@mui/material';

type Props = BoxProps;

export const Box = (props: Props) => <MuiBox {...props} />;