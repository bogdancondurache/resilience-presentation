int currentRetry = 0;

  for (;;)
  {
    try
    {
      await ExternalCallAsync();
      break;
    }
    catch (Exception ex)
    {
      currentRetry++;
      if (currentRetry > retryCount || !IsTransient(ex))
      {
        throw;
      }
    }

    await Task.Delay(delay);
  }