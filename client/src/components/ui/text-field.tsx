import { TextField as MuiTextField, TextFieldProps } from '@mui/material';

type Props = TextFieldProps;

export const TextField = (props: Props) => <MuiTextField {...props} />;
