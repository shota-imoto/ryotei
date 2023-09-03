import { ReactNode } from 'react';
import { Flex } from '../ui/flex';

type Props = {
  children: ReactNode;
};
export const Layout = ({ children }: Props) => {
  return <Flex width='100%'>{children}</Flex>;
};
