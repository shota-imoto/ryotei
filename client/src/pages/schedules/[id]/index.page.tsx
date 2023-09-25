import { Box } from '@/components/ui/box';
import { useQueryString } from './hooks/use-query-string';

const Schedule = () => {
  const id = useQueryString('id');
  return <Box>ID: {id}</Box>;
};

export default Schedule;
