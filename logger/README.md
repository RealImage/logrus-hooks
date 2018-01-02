## Logger

logger is a minimal hook for logrus that adds a `_logger` entry to indicate source of the log.

Its just a modification of a map field so subsequent `logger.Hook.Add()` hooks will be an overwrite.
