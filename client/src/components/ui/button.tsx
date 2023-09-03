import { ButtonProps, Button as MuiButton } from '@mui/material';

type Props = ButtonProps;

export const Button = (props: Props) => <MuiButton {...props} />;
