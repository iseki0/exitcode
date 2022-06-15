package exitcode

const (
	// ExOK successful termination
	ExOK Code = 0
)
const (
	// ExUsage The command was used incorrectly, e.g., with the wrong number of arguments, a bad flag, a bad syntax in a parameter, or whatever.
	ExUsage Code = iota + 64
	// ExDataErr The input data was incorrect in some way. This should only be used for user's data and not system files.
	ExDataErr
	// ExNoInput An input file (not a system file) did not exist or was not readable. This could also include errors like "No message" to a mailer (if it cared to catch it).
	ExNoInput
	// ExNoUser The user specified did not exist. This might be used for mail addresses or remote logins.
	ExNoUser
	// ExNoHost The host specified did not exist. This is used in mail addresses or network requests.
	ExNoHost
	// ExUnavailable A service is unavailable. This can occur if a support program or file does not exist. This can also be used as a catchall message when something you wanted to do does not work, but you do not know why.
	ExUnavailable
	// ExSoftware An internal software error has been detected. This should be limited to non-operating system related errors as possible.
	ExSoftware
	// ExOsErr An operating system error has been detected. This is intended to be used for such things as "cannot fork", "cannot create pipe", or the like. It includes things like getuid returning a user that does not exist in the passwd file.
	ExOsErr
	// ExOsFile Some system file (e.g., /etc/passwd, /var/run/utx.active, etc.) does not exist, cannot be opened, or has some sort of error (e.g., syntax error).
	ExOsFile
	// ExCantCreat A (user specified) output file cannot be created.
	ExCantCreat
	// ExIOErr An error occurred while doing I/O on some file.
	ExIOErr
	// ExTempFail Temporary failure, indicating something that is not really an error. In sendmail, this means that a mailer (e.g.) could not create a connection, and the request should be reattempted later.
	ExTempFail
	// ExProtocol The remote system returned something that was "not possible" during a protocol exchange.
	ExProtocol
	// ExNoPerm You did not have sufficient permission to perform the operation. This is not intended for file system problems, which should use EX_NOINPUT or EX_CANTCREAT, but rather for higher level permis- sions.
	ExNoPerm
	// ExConfig Something was found in an unconfigured or misconfigured state.
	ExConfig
)
