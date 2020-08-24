using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace ConfigChangerNet.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class LogController : ControllerBase
    {
        private readonly ILogger<LogController> _logger;

        public LogController(ILogger<LogController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public ActionResult Get()
        {
            Console.WriteLine("-----------------------------------------");
            _logger.LogDebug("DBG log");
            _logger.LogInformation("INF log");
            _logger.LogWarning("WRN log");
            _logger.LogError("ERR log");
            _logger.LogCritical("CRI log");
            return Ok();
        }
    }
}
