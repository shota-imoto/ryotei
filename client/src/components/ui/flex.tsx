import { Box } from '@mui/material';
import { ComponentProps, forwardRef } from 'react';

export type Props = Omit<ComponentProps<typeof Box>, 'display'>;

export const Flex = forwardRef<any, Props>(function Flex(props, ref) {
  return <Box ref={ref} {...props} display="flex" />;
});