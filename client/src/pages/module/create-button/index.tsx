import { Button } from '@/components/ui/button';
import { useRouter } from 'next/router';
import { useCallback } from 'react';

export const CreateButton = () => {
  const router = useRouter();
  const handleSubmit = useCallback(() => {
    router.push('/schedules/create');
  }, [router]);

  return <Button onClick={handleSubmit}>スケジュールを作成</Button>;
};
