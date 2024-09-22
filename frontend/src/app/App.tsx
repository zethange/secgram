import './index.css';

import { Router } from '@solidjs/router';
import { routes } from '@/app/routes.ts';
import { getCurrentUser } from '@/shared/lib/server/user';
import { currentUser } from '@/entities/user';

export const App = () => {
  (async () => {
    try {
      const currUser = await getCurrentUser()
      if (currUser) {
        currentUser.set(currUser)
      }
    } catch (e) {
      console.info("user is not authentificated")
    }
  })()

  return <Router children={routes} />
}
